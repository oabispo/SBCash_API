package test

import (
	"database/sql"
	"fmt"
	"math"

	plainRepo "sbcash_api/repositories/plain"

	plainModel "sbcash_api/model/mySql/plainData"

	richModel "sbcash_api/model/mySql/richData"

	helper "github.com/oabispo/BAMySQLHelper"
)

func TestaTudo(db *sql.DB) {
	if data, err := plainModel.Usuario_GetByID(db, 3); err == nil {
		fmt.Printf("\n%v", data)
	}

	if data, err := plainModel.Usuario_GetAll(db); err == nil {
		printData(data)
	}

	if data, err := helper.GetIntValue(db, "select count(1) from usuario u"); err == nil {
		fmt.Printf("\n\nNumero de usuarios: %v", data)
		pages := int(math.Ceil(float64(data) / float64(5)))

		for i := 1; i <= pages; i++ {
			if data, err := plainModel.Usuario_GetAllPaged(db, 5, i); err == nil {
				printData(data)
			} else {
				panic(err.Error())
			}
		}
	}

	if data, err := plainModel.Cartao_GetByID(db, 125989); err == nil {
		fmt.Printf("\n%v", data)
	} else {
		fmt.Printf("\n%v", err)
	}

	if data, err := richModel.Estabelecimento_GetInfo(db); err == nil {
		fmt.Printf("\n%v", data)
		richModel.Estabelecimento_UpdateInfo(db, data.Nome, data.RazaoSocial, data.CNPJ, data.Endereco, data.Cidade, data.Data_Instalacao, data.Hora_Abre, data.Hora_Fecha, data.Tempo_login_Expira, data.NumViasDescricao, data.Impressao_parcial, data.Idioma)
	}

	if data, err := richModel.Cupom_GetInfo(db); err == nil {
		fmt.Printf("\n%v", data)
		richModel.Cupom_UpdateInfo(db, data.LinhaA, data.LinhaB, data.DiasVigencia, data.Operacao, data.ValorMaxVenda)
	}

	if data, err := richModel.Cortesia_GetInfo(db); err == nil {
		fmt.Printf("\n%v", data)
		richModel.Cortesia_UpdateInfo(db, data.ValorMaximo, data.Vigencia, data.RegraCortesia, data.Valor1, data.Valor2, data.Valor3, data.Valor4, data.Valor5)
	}

	if data, err := richModel.Cortesia_Saldo(db); err == nil {
		fmt.Printf("\n%v", data)
	}

	if data, err := plainModel.Imposto_GetAll(db); err == nil {
		fmt.Printf("\n")
		for _, item := range data {
			fmt.Printf("\n%v", *item)
		}
	}

	if data, err := helper.GetIntValue(db, "select max(cod_imposto) from imposto"); err == nil {
		var cod_imposto = data

		if data, err := plainModel.Imposto_GetByID(db, cod_imposto); err == nil {
			fmt.Printf("\n%v", data)
		}
	}
	db.Close()
	fmt.Println("\n\nfim!")
}

func printData(items []*plainRepo.Usuario) {
	fmt.Printf("\n\n")
	for _, item := range items {
		fmt.Printf("\n%v", *item)
	}
}
