import '../../domain/entities/user.dart';

abstract class AuthState {
  AuthState();
}

class AuthInitial extends AuthState {
  AuthInitial();
}

class LoginLoading extends AuthState {
  LoginLoading();
}

class SignUpLoading extends AuthState {
  SignUpLoading();
}

class LoginSuccess extends AuthState {
  final String token;
  LoginSuccess({required this.token});
}

class SignUpSuccess extends AuthState {
  final User user;
  SignUpSuccess({required this.user});
}


class GetUserLoading extends AuthState {
  GetUserLoading();
}
class GetUserSuccess extends AuthState {
  final String name;
  GetUserSuccess({required this.name});
}
class AuthFailure extends AuthState {
  AuthFailure();
}