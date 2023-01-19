app.get('/auth', (req, res) => {
  const token = req.body.token;

  if (!token) {
    res.json({
      login: false,
      data: 'error'
    });
  return;
  }

  const decode = jwt.verify(token, 'secret');
  res.json({
      login: true,
      data: decode
  });
});