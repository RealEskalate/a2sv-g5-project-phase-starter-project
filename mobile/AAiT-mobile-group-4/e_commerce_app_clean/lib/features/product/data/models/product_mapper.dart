import '../../domain/entities/product_entity.dart';
import 'product_model.dart';

extension ProductMapper on ProductEntity {
  ProductModel toProductModel() {
    return ProductModel(
      id: id,
      name: name,
      description: description,
      price: price,
      imageUrl: imageUrl,
    );
  }
}


