# Write your MySQL query statement below
select c1.id, c1.movie, c1.description,c1.rating from cinema as c1 join
cinema as c2 on c1.id = c2.id where c1.id % 2 != 0 and c1.description != 'boring'
order by c1.rating desc;