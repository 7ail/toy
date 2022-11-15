package repository

type manager struct {
	fallback source
}

func Manager(fallback source) *manager {
	return &manager{
		fallback: fallback,
	}
}
