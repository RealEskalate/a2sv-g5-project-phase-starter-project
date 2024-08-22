import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../../data/models/product_models.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';

class AddProductUseCase  implements UseCase<String, ProductEntity> {
  final ProductRepository abstractProductRepository;

  AddProductUseCase(this.abstractProductRepository);

  @override
  Future<Either<Failure, String>> call(ProductEntity newProduct)async {
    return await abstractProductRepository.addProduct(newProduct); 

  }
}

class ProductParams extends Equatable {
  final String name;
  final String description;
  final double price;
  final String imagePath;
  const ProductParams(
      {required this.name, required this.description, required this.price,required this.imagePath});

  @override
  List<Object?> get props => [id, name, description, price,imagePath];
}

class AddProductParams  {
  final ProductEntity product;
  const AddProductParams({required this.product});
}
