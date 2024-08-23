import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  const Failure();

  @override
  List<Object> get props => [];
}

class ServerFailure extends Failure {
  final String message;

  const ServerFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class ConnectionFailure extends Failure {
  final String message;

  const ConnectionFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class CacheFailure extends Failure {
  final String message;

  const CacheFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class AuthenticationFailure extends Failure {
  final String message;

  const AuthenticationFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class AuthorizationFailure extends Failure {
  final String message;

  const AuthorizationFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class InvalidInputFailure extends Failure {
  final String message;

  const InvalidInputFailure({required this.message});

  @override
  List<Object> get props => [message];
}
class DatabaseFailure extends Failure {
  final String message;

  const DatabaseFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class PermissionFailure extends Failure {
  final String message;

  const PermissionFailure({required this.message});

  @override
  List<Object> get props => [message];
}

class LocationFailure extends Failure {
  final String message;

  const LocationFailure({required this.message});

  @override
  List<Object> get props => [message];
}
