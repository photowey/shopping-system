//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package auth

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &authenticator_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Authenticator{}
		},
	})
}

type authenticator_ struct {
	Check_ func(userID int64) bool
}

func (a *authenticator_) Check(userID int64) bool {
	return a.Check_(userID)
}

type AuthenticatorIOCInterface interface {
	Check(userID int64) bool
}

func GetAuthenticatorSingleton() (*Authenticator, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Authenticator)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Authenticator)
	return impl, nil
}

func GetAuthenticatorIOCInterfaceSingleton() (AuthenticatorIOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(Authenticator)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(AuthenticatorIOCInterface)
	return impl, nil
}
