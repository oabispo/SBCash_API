package plain

type Usuario struct {
	Cod_user      int    `json:"cod_user"`
	Cod_perfil    int    `json:"cod_perfil"`
	Nome          string `json:"nome"`
	Senha         string `json:"senha"`
	Status        int    `json:"status"`
	Aut_cortesia  bool   `json:"aut_cortesia"`
	Vis_pos_cx    bool   `json:"vis_pos_cx"`
}

func NewUsuario() *Usuario {
	return &Usuario{};
}

