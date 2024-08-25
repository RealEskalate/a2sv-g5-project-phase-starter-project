import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_entity.dart';

class ProductEntity extends Equatable {
  final String id;
  final String name;
  final String description;
  final int price;
  final String imageUrl;
  final UserEntity seller;

  const ProductEntity(
      {required this.id,
      required this.name,
      required this.description,
      required this.price,
      required this.imageUrl,
      required this.seller});

  @override
  List<Object?> get props => [
        id,
        name,
        description,
        price,
        imageUrl,
        seller,
      ];
}
