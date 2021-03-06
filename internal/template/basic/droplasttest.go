package basic

// DropLastTest is template to generate itself for different combination of data type.
func DropLastTest() string {
	return `
func TestDropLast<FTYPE>(t *testing.T) {
	list := []<TYPE>{1, 2, 3, 4, 5}
	expectedList := []<TYPE>{1, 2, 3, 4}
	actualList := DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{1, 2}
	expectedList = []<TYPE>{1}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{1}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}

// DropLastBoolTest is template to generate itself for different combination of data type.
func DropLastBoolTest() string {
	return `
func TestDropLast<FTYPE>(t *testing.T) {
	list := []<TYPE>{true, true, true, true, false}
	expectedList := []<TYPE>{true, true, true, true}
	actualList := DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{true, true}
	expectedList = []<TYPE>{true}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{true}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []<TYPE>{}
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []<TYPE>{}
	actualList = DropLast<FTYPE>(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLast<FTYPE> failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
`
}
