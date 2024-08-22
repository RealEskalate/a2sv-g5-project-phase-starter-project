part of 'sign_in_page_bloc.dart';

sealed class SignInPageEvent extends Equatable {
  const SignInPageEvent();

  @override
  List<Object> get props => [];
}

final class SignInPageButtonPressed extends SignInPageEvent {
  final String email;
  final String password;

  const SignInPageButtonPressed({required this.email, required this.password});

  @override
  List<Object> get props => [email, password];
}