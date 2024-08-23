

import 'package:equatable/equatable.dart';

abstract class LoginRegistrationState extends Equatable {

  const LoginRegistrationState ();
}
class InisialState extends LoginRegistrationState {
  const InisialState();

  @override
  List<Object?> get props => [];
}
class LoginRegistrationStarting extends LoginRegistrationState {
  final int showLogin;
  const LoginRegistrationStarting ({
    this.showLogin = 1
  });
  
  @override

  List<Object?> get props => [showLogin];
}

class OnLoading extends LoginRegistrationState {
  const OnLoading();
  
  @override
 
  List<Object?> get props => [];
}
class OnInputChange extends LoginRegistrationState {
  final String email;
  final String password;
  final String confirmPassword;
  final String fullName;
  final String newEmail;
  final String newPassword;
  final bool terms;

  const OnInputChange({
    required this.email,
    required this.password,
    required this.confirmPassword,
    required this.fullName,
    required this.newEmail,
    required this.newPassword,
    required this.terms,
 
    });
    
      @override
     
      List<Object?> get props => [email, password, confirmPassword, fullName];
  
}

class OnInputError extends LoginRegistrationState {
  final String error;
  const OnInputError({
    required this.error,
  });

  @override
  List<Object?> get props => [error];
}

class LoginSuccess extends LoginRegistrationState {
  final String email;
  final String name;
  

  const LoginSuccess({
    required this.email,
    required this.name,
   
   
  });

  @override
  List<Object?> get props => [email, name];
}

class RegistrationSuccess extends LoginRegistrationState {
  final bool success;

  const RegistrationSuccess({
    required this.success,
  });

  @override
  List<Object?> get props => [success];
}

class OnErrorState extends LoginRegistrationState {
  final String error;
  final bool email;
  final bool password;
  final bool fullName;
  final bool confirmPassword;
  final bool newEmail;
  final bool newPassword;
  final bool terms;
  const OnErrorState({
    required this.error,
    this.email = false,
    this.password = false,
    this.fullName = false,
    this.confirmPassword = false,
    this.newEmail = false,
    this.newPassword = false,
    this.terms = false,
  });

  @override
  List<Object?> get props => [error];
}