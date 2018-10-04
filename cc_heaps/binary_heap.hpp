
#pragma once

#include <vector>
#include <iostream>

namespace myheaps
{

template <typename DataType, typename KeyType = int, typename HandleType = int>
struct HeapNode
{
    KeyType key;      // i.e. priority
    HandleType index; // indicate index of this node in storage, start by 0

    DataType data; // external data

    HeapNode(KeyType k, DataType d) : key(k), data(d) {}
};

template <typename DataType, typename KeyType, typename HandleType>
std::ostream &operator<<(std::ostream &out, const HeapNode<DataType, KeyType, HandleType> &a)
{
    out << "{" << a.key << "," << a.index << "," << a.data << "}";
    return out;
}

template <typename DataType, typename KeyType, typename HandleType>
bool operator<(const HeapNode<DataType, KeyType, HandleType> &a, const HeapNode<DataType, KeyType, HandleType> &b)
{
    return a.key < b.key;
}

template <typename DataType, typename KeyType = int, typename HandleType = int>
class BinaryHeap
{
  protected:
    using NodeType = HeapNode<DataType, KeyType, HandleType>;

  public:
    bool empty() const { return storage_.empty(); }
    int size() const { return storage_.size(); }
    int capacity() const { return storage_.capacity(); }
    int heap_size() const { return heap_size_; }
    const std::vector<NodeType> &storage() const { return storage_; };

  public:
    // NOTE: `n` in below functions are all started by 1, not 0.
    bool IsMaxHeap(int n = 1) const
    {
        if (n > heap_size_)
        {
            return true; // exit rescurse
        }

        int left = leftChild(n);
        int right = rightChild(n);

        if (left <= heap_size_ && storage_[n - 1] < storage_[left - 1])
        {
            return false;
        }
        if (right <= heap_size_ && storage_[n - 1] < storage_[right - 1])
        {
            return false;
        }

        return IsMaxHeap(left) && IsMaxHeap(right);
    }

  protected:
    // NOTE: `n` in below functions are all started by 1, not 0.
    int leftChild(int n) const { return 2 * n; }
    int rightChild(int n) const { return 2 * n + 1; }
    int parent(int n) const { return n / 2; }

  protected:
    void Swap(int i, int j)
    {
        HeapNode<DataType> tmp = storage_[i];
        storage_[i] = storage_[j];
        storage_[j] = tmp;

        storage_[i].index = i;
        storage_[j].index = j;
    }

  protected:
    std::vector<NodeType> storage_;
    int heap_size_ = {0};
};

template <typename DataType, typename KeyType>
std::ostream &operator<<(std::ostream &out, const BinaryHeap<DataType, KeyType> &a)
{
    out << "{" << a.heap_size() << ", {";
    int i = 0;
    for (auto &v : a.storage())
    {
        if (i != 0)
        {
            out << ", ";
        }
        out << v;
        i++;
    }
    out << "}}";
    return out;
}

} // namespace myheaps
