

import 'package:equatable/equatable.dart';

abstract class BottumState  extends Equatable{}

class IntialState extends BottumState {
  @override
  List<Object?> get props => [];
}

class OnButtonActivate extends BottumState {
  final bool isActivate;

  OnButtonActivate ({
    this.isActivate = false
  });

  @override
  List<Object?> get props => [isActivate];
}

class AddErrorState extends BottumState {
  final String messages;

  AddErrorState ({
    required this.messages
  });

  @override
  List<Object ?> get props => [messages];
}




class AddLoadingState extends BottumState {
  @override
  List<Object?> get props => [];
}
class SuccessAddProduct extends BottumState {
  final bool add;

  SuccessAddProduct ({
    required this.add
  });

  @override

  List<Object?> get props => [add];
}