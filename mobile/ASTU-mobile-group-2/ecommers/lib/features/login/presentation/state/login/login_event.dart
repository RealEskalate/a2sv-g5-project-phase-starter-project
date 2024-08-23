

abstract class LoginEvent {}


class LoginRequest extends LoginEvent {
  final String email;
  final String password;

  LoginRequest({
    required this.email,
    required this.password
  });
}