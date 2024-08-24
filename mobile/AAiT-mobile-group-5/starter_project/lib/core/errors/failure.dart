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

class UnexpectedFailure extends Failure {
  final String message;

  const UnexpectedFailure({required this.message});

  @override
  List<Object> get props => [message];
}
