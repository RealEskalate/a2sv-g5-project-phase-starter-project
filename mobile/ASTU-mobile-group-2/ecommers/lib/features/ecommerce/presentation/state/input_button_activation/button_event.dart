
import 'dart:io';

import 'package:equatable/equatable.dart';

abstract class ButtonEvent  extends Equatable{}
class InsertInput extends ButtonEvent {
  final String name;
  final String price;
  final String description;
  final String category;
  final File? image;
  final String tag;
  final int type;
  final String id;
  InsertInput({
    this.name = '',
    this.price = '',
    this.description = '',
    this.category = '',
    this.image,
    this.tag = '',
    this.id = '',
    required this.type

  });
  @override
  List<Object ?> get props => [name,price,description,category,image];
}

class AddProductEvent extends ButtonEvent {
  AddProductEvent();
  @override
  List<Object ?> get props => [];
}

class UpdateProductEvent extends ButtonEvent {
  UpdateProductEvent();
  @override
  List<Object ?> get props => [];
}


