-- GetBookByUserId
SELECT book_id, book_name
FROM books
WHERE user_id = ?;
