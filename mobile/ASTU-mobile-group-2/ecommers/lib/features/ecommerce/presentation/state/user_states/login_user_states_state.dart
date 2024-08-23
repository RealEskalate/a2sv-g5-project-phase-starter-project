


import 'package:equatable/equatable.dart';

abstract class LoginUserStates extends Equatable{}


class LeftUserStates extends LoginUserStates {
  @override
  List<Object?> get props => [];
}

class UsernNameOfUser extends LoginUserStates {
  final String name;
  UsernNameOfUser({required this.name});
  @override
  List<Object?> get props => [];
}

class LogedOutUserStates extends LoginUserStates {
  @override
  List<Object?> get props => [];
}

class ProfileDetailState extends LoginUserStates {
  final String name;
  final String email;
  ProfileDetailState({required this.name,required this.email});
  @override
  List<Object?> get props => [];
}