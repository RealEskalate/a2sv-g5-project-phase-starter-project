import 'package:equatable/equatable.dart';

import '../../data/models/product_model.dart';

class ProductEntity extends Equatable {
  final String id;
  final String name;
  final String description;
  final double price;
  final String imageUrl;

  const ProductEntity({
    required this.id,
    required this.name,
    required this.description,
    required this.price,
    required this.imageUrl,
  });

  ProductModel toModel() => ProductModel(
        id: id,
        name: name,
        description: description,
        price: price,
        imageUrl: imageUrl,
      );

  @override
  List<Object?> get props => [
        id,
        name,
        description,
        price,
        imageUrl,
      ];
}
