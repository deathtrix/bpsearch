func testAVL() {
	// Add to AVL tree
	// var root *avl.Node
	// keys := []string{"2", "6", "1", "3", "5", "7", "16", "15", "14", "13", "12", "11", "8", "9", "10"}
	// for _, key := range keys {
	// 	root = avl.Insert(root, key)
	// }

	// avl.InOrder(root, func(node *avl.Node) {
	// 	fmt.Println(node.Key, node.Height)
	// })

	// var f = avl.Find(root, "7")
	// fmt.Printf("%+v\n", f)

	tree := avl.NewWithStringComparator() // empty(keys are of type int)

	tree.Put("1", "xfdsf")
	tree.Put("2", "fdsfb")
	tree.Put("1", "fdsfa")
	tree.Put("3", "cfdsf")
	tree.Put("4", "dfsf")
	tree.Put("5", "fse")
	tree.Put("6", "sssf")

	fmt.Println(tree)

	val, _ := tree.Get("4")
	fmt.Println(val)
	json, _ := tree.ToJSON()
	fmt.Printf("%+v\n", json)

	tree2 := avl.NewWithStringComparator()
	tree2.FromJSON(json)
	fmt.Println(tree2)

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	// dec := gob.NewDecoder(&b)

	err := enc.Encode(json)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Println(b)
	// output := base64.StdEncoding.EncodeToString(b.Bytes())
	err = ioutil.WriteFile("out", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// tree2 := avl.NewWithStringComparator()
	// err = dec.Decode(&tree2)
	// if err != nil {
	// 	log.Fatal("decode error 1:", err)
	// }
	// fmt.Println(tree2)
}
