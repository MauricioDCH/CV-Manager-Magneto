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

type EncryptedKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"column:key"`
}

type UserService interface {
	LoginUser(email, password string) (models.User, error)
}

type userService struct {
	db *gorm.DB
	key []byte
}

func init(){
	if err := godotenv.Load(); err != nil{
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

	masterKeyBase64 := os.Getenv("MASTER_KEY")
	masterKey, err := decodeBase64(masterKeyBase64)
	if err != nil {
		return nil, err
	}
	if len(masterKey) != keySize {
		return nil, errors.New("invalid key size")
	}

	encryptedKey, err := getEncryptedKey(db)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var key []byte
	if encryptedKey != "" {
		key, err = decrypt(encryptedKey, masterKey)
		if err != nil {
			return nil, err
		}
	} else {
		key, err = generateRandomKey()
		if err != nil {
			return nil, err
		}

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

func (s *userService) LoginUser(email, password string) (models.User, error) {
	var user models.User
	result := s.db.Where("correo = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("usuario no encontrado")
	}

	decryptedPassword, err := decrypt(user.Contraseña, s.key)
	if err != nil {
		return models.User{}, errors.New("error al desencriptar la contraseña")
	}

	if string(decryptedPassword) != password {
		return models.User{}, errors.New("contraseña incorrecta")
	}

	return user, nil
}