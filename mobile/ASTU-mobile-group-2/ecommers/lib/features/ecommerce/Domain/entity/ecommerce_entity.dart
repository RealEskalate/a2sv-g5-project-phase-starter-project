



import 'package:equatable/equatable.dart';

import '../../Data/model/ecommerce_model.dart';

class EcommerceEntity extends Equatable{
  const EcommerceEntity ({
    required this.id,
    required this.name,
    required this.description,
    required this.imageUrl,
    required this.price
  });
  final String id;
  final String name;
  final String description;
  final String imageUrl;
  final double price;
  EcommerceModel toModel() => EcommerceModel(
    id:id,
    name: name,
    description: description,
    imageUrl: imageUrl,
    price: price
  );
  @override
  List<Object?> get props => [
    id,
    name,
    description,
    imageUrl,
    price,
  ];

}