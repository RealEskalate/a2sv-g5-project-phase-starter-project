import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_entity.dart';
import '../../data/models/product_model.dart';

class ProductEntity extends Equatable {
  final String id;
  final String name;
  final String description;
  final double price;
  final String imageUrl;
  final UserEntity seller;

  const ProductEntity({
    required this.id,
    required this.name,
    required this.description,
    required this.price,
    required this.imageUrl,
    required this.seller,
  });

  ProductModel toModel() => ProductModel(
      id: id,
      name: name,
      description: description,
      price: price,
      imageUrl: imageUrl,
      seller: seller);

  @override
  List<Object?> get props => [
        id,
        name,
        description,
        price,
        imageUrl,
      ];
}
