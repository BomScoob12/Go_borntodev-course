## 2 type of route
- **Static Route**: The URL pattern is fixed and does not change.
- **Dynamic Route**: The URL pattern is variable and can change.

### example of static route
```javascript
app.get('/about', (req, res) => {
  res.send('About Page');
});
```
### example of dynamic route
```javascript
app.get('/post/:id', (req, res) => {
  res.send('Post ' + req.params.id);
});
```
### example of dynamic route with multiple parameters
```javascript
app.get('/post/:id/:category', (req, res) => {
  res.send('Post ' + req.params.id + ' Category ' + req.params.category);
});
```
### example of dynamic route with optional parameters
```javascript
app.get('/post/:id/:category?', (req, res) => {
  res.send('Post ' + req.params.id + ' Category ' + req.params.category);
});
```