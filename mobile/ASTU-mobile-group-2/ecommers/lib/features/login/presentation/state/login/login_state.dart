

import 'package:equatable/equatable.dart';

abstract class LoginState {
  LoginState();
}

class LoginIntial extends LoginState {
  LoginIntial();
}

class LoginSFuiled extends LoginState with EquatableMixin{
  final String message;
  LoginSFuiled({
    required this.message
  });
  
  @override
  List<Object?> get props => [message];
}

class LoginSuccess extends LoginState with EquatableMixin{
  final String message;
  LoginSuccess ({
    required this.message,
  });
  @override
  List<Object?> get props => [message];
}