import 'package:equatable/equatable.dart';

class SignedInEntity extends Equatable {
  final String accessToken;

  const SignedInEntity({required this.accessToken});

  @override
  List<Object?> get props => [accessToken];
}
