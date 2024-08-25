
import 'package:equatable/equatable.dart';

abstract class ImageEvent extends Equatable{}


class SelectImageEvent extends ImageEvent {
 
   SelectImageEvent();
  
  @override
  List<Object ?> get props => [];
}

class InputIntialEvent extends ImageEvent {
  @override
  List<Object?> get props => [];
}