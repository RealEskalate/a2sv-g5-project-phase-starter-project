


import 'package:equatable/equatable.dart';

abstract class LoginUserStatesEvent  extends Equatable{}

class GetUserNameEvent extends LoginUserStatesEvent {

  @override
  List<Object?> get props => [];
}

class LogedOutUserStatesEvent extends LoginUserStatesEvent {
  @override
  List<Object?> get props => [];
}

class ProfileDetail extends LoginUserStatesEvent {
  @override
  List<Object?> get props => [];
}