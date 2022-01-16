# Building Go Web Applications and Microservices Using Gin

In this tutorial, you will learn how to build traditional web applications and microservices in Go using the Gin framework. Gin is a framework that reduces boilerplate code that would normally go into building these applications. It also lends itself very well to creating reusable and extensible pieces of code.

## Requirements

- `Go`
- `Git`
- `Gin`

## Application Functionality

The application we’ll build is a simple article manager. This application should:

- Let users register with a username and a password (non-logged in users only),
- Let users login with a username and a password (non-logged in users only),
- Let users log out (logged in users only),
- Let users create new articles (logged in users only),
- Display the list of all articles on the home page (for all users), and
- Display a single article on its own page (for all users).

In addition to this functionality, the list of articles and a single article should be accessible in the HTML, JSON and XML formats.

This will allow us to illustrate how Gin can be used to design traditional web applications, API servers, and microservices.

To achieve this, we will make use of the following functionalities offered by Gin:

- Routing — to handle various URLs,
- Custom rendering — to handle the response format, and
- Middleware — to implement authentication.

We’ll also [write tests to validate](https://semaphoreci-com.translate.goog/community/tutorials/building-go-web-applications-and-microservices-using-gin?_x_tr_sl=en&_x_tr_tl=tr&_x_tr_hl=tr&_x_tr_pto=wapp) that all the features work as intended.

## Routing

In our application, we will:

- Serve the index page at route / (HTTP GET request),
- Group user-related routes under the /u route,
  - Serve the login page at /u/login (HTTP GET request),
  - Process the login credentials at /u/login (HTTP POST request),
  - Log out at /u/logout (HTTP GET request),
  - Serve the registration page at /u/register (HTTP GET request),
  - Process the registration information at /u/register (HTTP POST request) ,
- Group article related routes under the /article route,
  - Serve the article creation page at /article/create (HTTP GET request),
  - Process the submitted article at /article/create (HTTP POST request), and
  - Serve the article page at /article/view/:article_id (HTTP GET request). Take note of the :article_id part in this route. The : at the beginning indicates that this is a dynamic route. This means that :article_id can contain any value and Gin will make this value available in the route handler.

## Testing the Application

```Bash
# Retrieving the List of Articles in JSON Format
$curl -X GET -H "Accept: application/json" http://localhost:8080/
[{"id":1,"title":"Article 1","content":"Article 1 body"},{"id":2,"title":"Article 2","content":"Article 2 body"}]
# Retrieving an Article in XML Format
$curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1
<article><ID>1</ID><Title>Article 1</Title><Content>Article 1 body</Content></article>
# Run test
$go test -v
=== RUN   TestShowIndexPageUnauthenticated
[GIN] 2022/01/16 - 21:21:16 | 200 | 314.912µs || GET "/"
--- PASS: TestShowIndexPageUnauthenticated (0.00s)
=== RUN   TestArticleUnauthenticated
[GIN] 2022/01/16 - 21:21:16 | 200 | 104.032µs || GET "/article/view/1"
--- PASS: TestArticleUnauthenticated (0.00s)
=== RUN   TestArticleListJSON
[GIN] 2022/01/16 - 21:21:16 | 200 | 44.106µs || GET "/"
--- PASS: TestArticleListJSON (0.00s)
=== RUN   TestArticleXML
[GIN] 2022/01/16 - 21:21:16 | 200 | 48.521µs || GET "/article/view/1"
--- PASS: TestArticleXML (0.00s)
=== RUN   TestGetAllArticles
--- PASS: TestGetAllArticles (0.00s)
=== RUN   TestGetArticleByID
--- PASS: TestGetArticleByID (0.00s)
PASS
ok      goAPI   0.087s

```

### Reference

[Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)
