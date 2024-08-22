import 'package:equatable/equatable.dart';

import '../../domain/entity/auth_entity.dart';

abstract class AuthEvent extends Equatable {
  const AuthEvent();

  @override
  List<Object> get props => [];
}

class LoginEvent extends AuthEvent {
  final String email;
  final String password;

  const LoginEvent({
    required this.email,
    required this.password,
  });

  @override
  List<Object> get props => [email, password];
}

class RegisterEvent extends AuthEvent {
  final SentUserEntity senduserentity;

  RegisterEvent({required this.senduserentity});


  
  @override
  List<Object> get props => [senduserentity];
}

class GetUserProfileEvent extends AuthEvent {
  

   GetUserProfileEvent();

  @override
  List<Object> get props => [];
}

class LogoutEvent extends AuthEvent {
  const LogoutEvent();

  @override
  List<Object> get props => [];
}
