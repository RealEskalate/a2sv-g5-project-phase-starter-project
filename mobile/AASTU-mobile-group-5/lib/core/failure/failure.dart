abstract class Failure {
  final String message;
  Failure({required this.message});

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is Failure &&
          runtimeType == other.runtimeType &&
          message == other.message;

  @override
  int get hashCode => message.hashCode;
}

class ServerFailure extends Failure {
  ServerFailure({required super.message});
}

class CacheFailure extends Failure {
  CacheFailure({required super.message});
}

class NetworkFailure extends Failure {
  NetworkFailure({required super.message});

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is NetworkFailure &&
          runtimeType == other.runtimeType &&
          message == other.message;

  @override
  int get hashCode => message.hashCode;
}
