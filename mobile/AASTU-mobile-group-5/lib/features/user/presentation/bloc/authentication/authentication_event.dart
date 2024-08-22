part of 'authentication_bloc.dart';

sealed class AuthenticationEvent extends Equatable {
  const AuthenticationEvent();

  @override
  List<Object> get props => [];
}

class LoggedOut extends AuthenticationEvent {}

class CheckCurrentStatus extends AuthenticationEvent {}