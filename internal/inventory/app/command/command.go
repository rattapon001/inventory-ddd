package command

type InventoryCommand interface {
	execute(payload []byte) error
}
