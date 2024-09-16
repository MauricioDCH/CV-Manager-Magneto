package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"os"

	"CV_MANAGER/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const keySize = 32

// EncryptedKey defines the structure for storing encrypted keys in the database.
type EncryptedKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"column:key"`
}

// UserService defines the interface for user services.
type UserService interface {
	RegisterUser(name, email, password string) (models.User, error)
}

type userService struct {
	db  *gorm.DB
	key []byte // Encryption key
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func generateRandomKey() ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func saveEncryptedKey(db *gorm.DB, encryptedKey string) error {
	keyRecord := EncryptedKey{Key: encryptedKey}
	if err := db.Create(&keyRecord).Error; err != nil {
		return err
	}
	return nil
}

func getEncryptedKey(db *gorm.DB) (string, error) {
	var keyRecord EncryptedKey
	result := db.First(&keyRecord)
	if result.Error != nil {
		return "", result.Error
	}
	return keyRecord.Key, nil
}

func decodeBase64(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}

func decrypt(ciphertext string, key []byte) ([]byte, error) {
	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return ciphertextBytes, nil
}

func NewUserService(db *gorm.DB) (UserService, error) {
	// Retrieve the master key from environment variables
	masterKeyBase64 := os.Getenv("MASTER_KEY")
	masterKey, err := decodeBase64(masterKeyBase64)
	if err != nil {
		return nil, err
	}
	if len(masterKey) != keySize {
		return nil, errors.New("invalid key size")
	}

	// Try to retrieve the encrypted key from the database
	encryptedKey, err := getEncryptedKey(db)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var key []byte
	if encryptedKey != "" {
		// Decrypt the key
		key, err = decrypt(encryptedKey, masterKey)
		if err != nil {
			return nil, err
		}
	} else {
		// Generate a new key and save it
		key, err = generateRandomKey()
		if err != nil {
			return nil, err
		}

		// Encrypt and save the new key
		encryptedKey, err = encrypt(string(key), masterKey)
		if err != nil {
			return nil, err
		}
		if err := saveEncryptedKey(db, encryptedKey); err != nil {
			return nil, err
		}
	}

	return &userService{db: db, key: key}, nil
}

func (s *userService) RegisterUser(name, email, password string) (models.User, error) {
	if name == "" {
		return models.User{}, errors.New("el nombre no puede estar vacío")
	}
	if password == "" {
		return models.User{}, errors.New("la contraseña no puede estar vacía")
	}

	var existingUser models.User
	result := s.db.Where("correo = ?", email).First(&existingUser)
	if result.Error == nil {
		return models.User{}, errors.New("email ya registrado")
	}

	encryptedPassword, err := encrypt(password, s.key)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Nombre:     name,
		Correo:     email,
		Contraseña: encryptedPassword,
	}

	if err := s.db.Create(&user).Error; err != nil {
		log.Println("Error al registrar el usuario:", err)
		return models.User{}, errors.New("error al registrar el usuario")
	}

	return user, nil
}
