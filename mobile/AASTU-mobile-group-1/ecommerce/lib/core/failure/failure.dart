import 'package:equatable/equatable.dart';

// ignore: must_be_immutable
class Failure extends Equatable {
  String message;
  Failure({required this.message});
  @override
  // TODO: implement props
  List<Object?> get props => [message];
}
