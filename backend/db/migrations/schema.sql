create database if not exists my_vocabulary_book;
use my_vocabulary_book;

-- table
create table user (
  user_id int auto_increment,
  user_name varchar(50) not null,
  password varchar(20),
  create_at timestamp default current_timestamp,
  primary key(user_id)
);

create table vocabulary_book (
  book_id int auto_increment,
  user_id int,
  book_name varchar(10) not null,
  primary key(book_id),
  foreign key(user_id)
    references user(user_id) on delete cascade
);

create table word (
  word_id int auto_increment,
  book_id int not null,
  word varchar(50),
  translated_word varchar(50),
  create_at timestamp default current_timestamp,
  primary key(word_id),
  foreign key(book_id)
    references vocabulary_book(book_id) on delete cascade
);
