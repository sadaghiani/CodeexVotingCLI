package service

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// CreateAccount generate new ethereum address
func (s *Eth) CreateAccount(password string) (string, error) {

	if password == "" {
		return "", fmt.Errorf("password is required")
	}

	account, err := s.keystore.NewAccount(password)
	if err != nil {
		log.Printf("Error : %v", err)
		return "", err
	}
	return account.Address.Hex(), nil
}

// AccountsList return list of public address
func (s *Eth) AccountsList() ([]string, error) {

	var accountList []string
	for _, val := range s.keystore.Accounts() {
		accountList = append(accountList, val.Address.Hex())
	}

	if len(accountList) == 0 {
		return nil, fmt.Errorf("not key stored")
	}

	return accountList, nil
}

// NewAccountFromPrivateKey create new account from private key in keystore
func (s *Eth) NewAccountFromPrivateKey(pk string, passphrase string) (bool, error) {

	defer s.client.Close()

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return false, err
	}

	s.keystore.ImportECDSA(privateKey, passphrase)

	return true, nil
}
