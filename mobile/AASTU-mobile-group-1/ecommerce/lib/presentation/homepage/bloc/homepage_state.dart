
part of 'homepage_bloc.dart';
enum HomepageStatus { initial, loading, loaded, error }

class HomepageState extends Equatable {
  final HomepageStatus status;
  final List<ProductEntity> products;

  const HomepageState({
    this.status = HomepageStatus.initial,
    this.products = const <ProductEntity>[],
  });

  HomepageState copyWith({
    HomepageStatus? status,
    List<ProductEntity>? products,
    Failure? failure,
  }) {
    return HomepageState(
      status: status ?? this.status,
      products: products ?? this.products,
    );
  }

  @override
  List<Object> get props => [status, products];
}