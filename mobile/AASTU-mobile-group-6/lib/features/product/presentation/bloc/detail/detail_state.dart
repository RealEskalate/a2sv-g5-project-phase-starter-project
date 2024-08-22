import 'package:equatable/equatable.dart';




// part of 'home_bloc.dart';

abstract class DetailState extends Equatable {
  const DetailState();
}

class DetailLoading extends DetailState {
  @override
  List<Object> get props => [];
}

class DetailLoaded extends DetailState {
  final String message;

  const DetailLoaded(this.message);
  @override
  List<Object> get props => [message];

}
class DetailFailure extends DetailState {
  final String error;

  const DetailFailure(this.error);

  @override
  List<Object> get props => [error];
}
