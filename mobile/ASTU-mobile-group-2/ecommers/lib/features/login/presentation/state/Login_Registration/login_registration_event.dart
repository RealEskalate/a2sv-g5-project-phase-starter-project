abstract class LoginRegistrationEvent {}


class OnLogin extends LoginRegistrationEvent{
  OnLogin();
}

class OnRegistration extends LoginRegistrationEvent{}

class OnForgot extends LoginRegistrationEvent{}

class OnInputChangeEvent extends LoginRegistrationEvent{
  final String email;
  final String password;
  final String confirmPassword;
  final String fullName;
  final String type;
  final String newEmail;
  final String newPassword;
  final bool terms;


  OnInputChangeEvent({
    this.email = '',
    this.password = '',
    this.confirmPassword = '',
    this.fullName = '',
    required this.type,
    this.newEmail = '',
    this.newPassword = '',
    this.terms = false,

    });
}

class LoginButtonPressed extends LoginRegistrationEvent{
  LoginButtonPressed();
}


class RegisterButtonPressed extends LoginRegistrationEvent{
  RegisterButtonPressed();
}