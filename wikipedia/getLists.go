package wikipedia

func (s *session) GetLists() []ListInfo {
	result := make([]ListInfo, 0, len(s.Lists))

	for _, list := range s.Lists {
		result = append(result, list.ListInfo)
	}

	return result
}
