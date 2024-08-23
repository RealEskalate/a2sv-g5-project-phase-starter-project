import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  final String message;

  const Failure(this.message);

  @override
  List<Object> get props => [message];
}

class ServerFailure extends Failure {
  const ServerFailure(super.message);
}

class ConnectionFailure extends Failure {
  const ConnectionFailure()
      : super('No Internet connection. Please check your connection.');
}

class DatabaseFailure extends Failure {
  const DatabaseFailure(super.message);
}

class WrongPasswordOrEmailFailure extends Failure {
  const WrongPasswordOrEmailFailure(super.message);
}

class InvalidEmailFailure extends Failure {
  const InvalidEmailFailure(super.message);
}

class AlreadyExistEmailFailure extends Failure {
  const AlreadyExistEmailFailure(super.message);
}

class CacheFailure extends Failure {
  const CacheFailure(super.message);
}

class UnkownFailure extends Failure {
  const UnkownFailure()
      : super('An unknown error occurred. Please try again later.');
}
