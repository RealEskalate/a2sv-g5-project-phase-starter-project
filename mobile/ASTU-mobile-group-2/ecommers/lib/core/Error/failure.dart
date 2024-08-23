

import 'package:equatable/equatable.dart';



abstract class Failure extends Equatable {
  final String message;
  const Failure({
    required this.message
  });
  @override
  List<Object> get props => [message];
}


class ServerFailure extends Failure {
  const ServerFailure ({
    required super.message
  });
}

class ConnectionFailur extends Failure {

  const ConnectionFailur ({
    required super.message
  });

}
class CachException extends Failure {
  const CachException ({
    required super.message
  });
}