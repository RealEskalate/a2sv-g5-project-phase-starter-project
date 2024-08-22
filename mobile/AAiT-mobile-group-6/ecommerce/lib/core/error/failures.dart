import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  final String message;
  const Failure(this.message);
  @override
  List<Object> get props => [message];
}

class ServerFailure extends Failure {
  ServerFailure(String message) : super(message);
  @override
  List<Object> get props => [message];
}

class CacheFailure extends Failure {
  CacheFailure(String message) : super(message);
  @override
  List<Object> get props => [message];
}

class NetworkFailure extends Failure {
  NetworkFailure(String message) : super(message);
  @override
  List<Object> get props => [message];
}
