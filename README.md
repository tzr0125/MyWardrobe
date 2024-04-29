# MyWardrobe

Self-hosted wbsite to manage your clothes

setup postgresql:

docker run --rm --name mywardrobe_db_debug -v mywardrobe_postgres:/var/lib/postgresql/data -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_USER=root -e POSTGRES_DB=mywardrobe -p 6666:5432 -d postgres
