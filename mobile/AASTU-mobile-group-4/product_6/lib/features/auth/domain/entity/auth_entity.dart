import 'package:equatable/equatable.dart';

class AuthEntity extends Equatable {
  final String email;
  final String password;

  AuthEntity({required this.email, required this.password});

  @override
  List<Object?> get props => [email, password];
}

class UserEntity extends Equatable {
  final String id;
  final String name;
  final String email;

  UserEntity({required this.id, required this.name, required this.email});

  @override
  List<Object?> get props => [id, name, email];
}


class SentUserEntity extends Equatable {

  final String name;
  final String email;
  final String password;

  SentUserEntity({ required this.name, required this.email,required this.password});

  @override
  List<Object?> get props => [ name, email];
}



class AuthResponseEntity extends Equatable {
  final String accessToken;

  AuthResponseEntity({required this.accessToken});

  @override
  List<Object?> get props => [accessToken];
}