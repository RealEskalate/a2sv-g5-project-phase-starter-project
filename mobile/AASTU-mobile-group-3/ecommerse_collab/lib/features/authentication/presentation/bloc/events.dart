

class UserEvent {}

class RegisterUserEvent extends UserEvent {
  String email;
  String username;
  String password;

  RegisterUserEvent({required this.email, required this.password, required this.username});
}

class LogInEvent extends UserEvent{
  String email;
  String password;
  LogInEvent({required this.email, required this.password});

}

class LogOutEvent extends UserEvent {}