package model

type Seq struct {
	SeqName string `bson:"seq_name"`
	SeqVal  int    `bson:"seq_val"`
}
