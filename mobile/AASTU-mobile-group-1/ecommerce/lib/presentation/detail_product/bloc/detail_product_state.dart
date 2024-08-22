part of 'detail_product_bloc.dart';

enum DetailProductStatus { initial, loading, loaded, error }

class DetailProductState extends Equatable{
  final DetailProductStatus status;
  final ProductEntity product;

  DetailProductState({
    this.status = DetailProductStatus.initial,
    required this.product,
  });
  @override
  // TODO: implement props
  List<Object?> get props => throw UnimplementedError();

}