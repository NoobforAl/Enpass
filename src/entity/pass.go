package entity

import (
	"context"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

type Pass struct {
	PassID    uint   `form:"passid" json:"passid"`
	UserID    uint   `form:"userid" json:"userid"`
	ServiceID uint   `form:"serviceid" json:"serviceid" binding:"required"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Note      string `form:"note" json:"note"`
}

func (p *Pass) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, p)
}

func (p *Pass) FindPassword(ctx context.Context, stor contract.Stor,
	dePass string, decrypt bool) error {
	pass, err := stor.GetPassword(ctx, p.PassID)
	if err != nil {
		return err
	}

	if decrypt {
		if err = pass.Values.DecryptValues(dePass); err != nil {
			return err
		}
	}

	p.PassID = pass.ID
	p.UserID = pass.UserID
	p.ServiceID = pass.ServiceID
	p.UserName = pass.UserName.String()
	p.Password = pass.Password.String()
	p.Note = pass.Note.String()
	return nil
}

func (p *Pass) GetAllPassword(ctx context.Context, stor contract.Stor,
	dePass string, decrypt bool) ([]Pass, error) {

	allPassword, err := stor.GetManyPassword(ctx)
	if err != nil {
		return nil, err
	}

	var allPass []Pass
	for _, v := range allPassword {
		if decrypt {
			if err = v.Values.DecryptValues(dePass); err != nil {
				return nil, err
			}
		}
		allPass = append(allPass, Pass{
			PassID:    v.ID,
			UserID:    v.UserID,
			ServiceID: v.ServiceID,
			UserName:  v.UserName.String(),
			Password:  v.Password.String(),
			Note:      v.Note.String(),
		})
	}
	return allPass, nil
}

func (p *Pass) UpdatePass(ctx context.Context, stor contract.Stor, dePass string) error {
	_, err := stor.GetPassword(ctx, p.PassID)
	if err != nil {
		return err
	}

	pass := stor.NewPassword(p.PassID, p.ServiceID, p.UserID,
		p.UserName, p.Password, p.Note, "")

	if err = pass.Values.EncryptValues(dePass); err != nil {
		return err
	}
	return stor.UpdatePassword(ctx, *pass)
}

func (p *Pass) CreatePass(ctx context.Context, stor contract.Stor, enPass string) error {
	pass := stor.NewPassword(p.PassID, p.ServiceID, p.UserID,
		p.UserName, p.Password, p.Note, "")

	if err := pass.Values.EncryptValues(enPass); err != nil {
		return err
	}

	return stor.InsertPassword(ctx, *pass)
}

func (p *Pass) DeletePass(ctx context.Context, stor contract.Stor) error {
	return stor.DeletePassword(ctx, p.PassID)
}
