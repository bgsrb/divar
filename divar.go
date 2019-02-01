package divar

//GetPostList ...
func (c *Client) GetPostList(req Request) ([]Post, error) {
	res := new(Response)
	if err := c.execute(req, res); err != nil {
		return []Post{}, err
	}
	return res.Result.PostList, nil
}
