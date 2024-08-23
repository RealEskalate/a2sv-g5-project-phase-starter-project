import 'dart:io';

import 'package:equatable/equatable.dart';

abstract class ProductEvents extends Equatable {
  @override
  List<Object?> get props => [];
}

class LoadAllProductEvents extends ProductEvents {}

class GetSingleProductEvents extends ProductEvents {
  final String id;
  GetSingleProductEvents({required this.id});
}

class UpdateProductEvent extends ProductEvents {
  final String id;
  final String name;
  final String description;
  final String price;

  UpdateProductEvent({
    required this.id,
    required this.name,
    required this.description,
    required this.price,
  });
}

class DeleteProductEvent extends ProductEvents {
  final String id;

  DeleteProductEvent({required this.id});
}

class RefreshEvent extends ProductEvents {}

class InsertProductEvent extends ProductEvents {
  final String name;
  final String description;
  final String price;
  final File imageUrl;

  InsertProductEvent({
    required this.name,
    required this.description,
    required this.price,
    required this.imageUrl,
  });
}
