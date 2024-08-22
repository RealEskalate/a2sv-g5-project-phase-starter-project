import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  final String message;
  const Failure(this.message);

  @override
  List<Object?> get props => [message];
}

class ServerFailure extends Failure {
  const ServerFailure(String message) : super(message);
}

class ConnectionFailure extends Failure {
  const ConnectionFailure(String message) : super(message);
}

class CacheFailure extends Failure {
  const CacheFailure(String message) : super(message);
}

class RandomFailure extends Failure {
  const RandomFailure(String message) : super(message);
}

class UnauthorizedFailure extends Failure{
  UnauthorizedFailure(super.message);
}

class UserAlreadyExistsFailure extends Failure{
  UserAlreadyExistsFailure(super.message);
}