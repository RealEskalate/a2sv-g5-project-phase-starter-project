part of 'sign_up_page_bloc.dart';

sealed class SignUpPageEvent extends Equatable {
  const SignUpPageEvent();

  @override
  List<Object> get props => [];
}

final class SignUpPageButtonPressed extends SignUpPageEvent {
  final String name;
  final String email;
  final String password;
  final String confirmPassword;

  const SignUpPageButtonPressed({
    required this.name,
    required this.email,
    required this.password,
    required this.confirmPassword,
  });

  @override
  List<Object> get props => [name, email, password, confirmPassword];
}