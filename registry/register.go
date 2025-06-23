package registry

import (
	"errors"
	kreg "github.com/go-kratos/kratos/v2/registry"
)

type Registry interface {
	Register(registry *RegistryConfig) kreg.Registrar
	Discovery(registry *RegistryConfig) kreg.Discovery
}

func NewReg(reg *RegistryConfig) (kreg.Registrar, error) {
	if reg.Nacos.Enable {
		return NewNacos().Register(reg), nil
	}
	return nil, errors.New("注册文件配置每启用或者没找到")
}
func NewDis(reg *RegistryConfig) (kreg.Discovery, error) {
	if reg.Nacos.Enable {
		return NewNacos().Discovery(reg), nil
	}
	return nil, errors.New("注册文件配置每启用或者没找到")
}

func NewRegDis(reg *RegistryConfig) (kreg.Registrar, kreg.Discovery, error) {
	newReg, err := NewReg(reg)
	if err != nil {
		return nil, nil, err
	}
	newDis, err := NewDis(reg)
	if err != nil {
		return nil, nil, err
	}
	return newReg, newDis, nil
}
