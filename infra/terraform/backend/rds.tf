resource "aws_db_instance" "psql_instance" {
  allocated_storage     = 5
  db_name               = "psql_gonext"
  multi_az              = false
  engine                = "postgres"
  engine_version        = "17.4"
  instance_class        = "db.t3.micro"
  username              = "postgres" //! to change
  password              = "postgres" //! to change
  skip_final_snapshot = true

  vpc_security_group_ids = [aws_security_group.db_sg.id]
}