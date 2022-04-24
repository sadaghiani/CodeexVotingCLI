package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewVoting returns the BalanceOf address in contract
func (s *Eth) NewVoting(from string, pass string, ID int64, minDeposit int64, endRegisterTime int64, endVotingTime int64) (string, error) {

	fromAddress := common.HexToAddress(from)

	defer func() {
		s.keystore.Lock(fromAddress)
		s.client.Close()
	}()

	fromAddressOnKeyStore, err := s.keystore.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return "", err
	}
	err = s.keystore.Unlock(fromAddressOnKeyStore, pass)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(s.keystore, fromAddressOnKeyStore, s.chainID)
	if err != nil {
		return "", err
	}

	tx, err := s.instance.RegisterVoting(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, big.NewInt(ID), big.NewInt(minDeposit), big.NewInt(endRegisterTime), big.NewInt(endVotingTime))
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// SubmitVote returns the BalanceOf address in contract
func (s *Eth) SubmitVote(from string, pass string, ID int64, name string) (string, error) {

	fromAddress := common.HexToAddress(from)

	defer func() {
		s.keystore.Lock(fromAddress)
		s.client.Close()
	}()

	fromAddressOnKeyStore, err := s.keystore.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return "", err
	}
	err = s.keystore.Unlock(fromAddressOnKeyStore, pass)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(s.keystore, fromAddressOnKeyStore, s.chainID)
	if err != nil {
		return "", err
	}

	nameByte := [32]byte{}
	copy(nameByte[:], []byte(name))

	tx, err := s.instance.SubmitVote(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, big.NewInt(ID), nameByte)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// ClearVote returns the BalanceOf address in contract
func (s *Eth) ClearVote(from string, pass string, ID int64, name string) (string, error) {

	fromAddress := common.HexToAddress(from)

	defer func() {
		s.keystore.Lock(fromAddress)
		s.client.Close()
	}()

	fromAddressOnKeyStore, err := s.keystore.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return "", err
	}
	err = s.keystore.Unlock(fromAddressOnKeyStore, pass)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(s.keystore, fromAddressOnKeyStore, s.chainID)
	if err != nil {
		return "", err
	}

	nameByte := [32]byte{}
	copy(nameByte[:], []byte(name))

	tx, err := s.instance.ClearVote(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, big.NewInt(ID), nameByte)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// Deposit returns the BalanceOf address in contract
func (s *Eth) Deposit(from string, pass string, ID int64, name string, value int64) (string, error) {

	fromAddress := common.HexToAddress(from)

	defer func() {
		s.keystore.Lock(fromAddress)
		s.client.Close()
	}()

	fromAddressOnKeyStore, err := s.keystore.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return "", err
	}
	err = s.keystore.Unlock(fromAddressOnKeyStore, pass)
	if err != nil {
		return "nil", err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(s.keystore, fromAddressOnKeyStore, s.chainID)
	if err != nil {
		return "", err
	}

	nameByte := [32]byte{}
	copy(nameByte[:], []byte(name))

	tx, err := s.instance.Deposit(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(value),
	}, big.NewInt(ID), nameByte)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

// Withdraw returns the BalanceOf address in contract
func (s *Eth) Withdraw(from string, pass string, ID int64) (string, error) {

	fromAddress := common.HexToAddress(from)

	defer func() {
		s.keystore.Lock(fromAddress)
		s.client.Close()
	}()

	fromAddressOnKeyStore, err := s.keystore.Find(accounts.Account{Address: fromAddress})
	if err != nil {
		return "", err
	}
	err = s.keystore.Unlock(fromAddressOnKeyStore, pass)
	if err != nil {
		return "nil", err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(s.keystore, fromAddressOnKeyStore, s.chainID)
	if err != nil {
		return "", err
	}

	tx, err := s.instance.Withdraw(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, big.NewInt(ID))
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

// GetListNameFromID returns the BalanceOf address in contract
func (s *Eth) GetListNameFromID(ID int64) ([]string, error) {

	defer func() {
		s.client.Close()
	}()

	names, err := s.instance.GetListNameFromID(&bind.CallOpts{}, big.NewInt(ID))
	if err != nil {
		return nil, err
	}

	var sName []string
	for _, v := range names {
		sName = append(sName, string(v[:]))
	}
	return sName, nil
}
