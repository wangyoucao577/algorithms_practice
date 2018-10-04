
#pragma once

#include "binary_heap.hpp"
#include <assert.h>

namespace myheaps
{

template <typename DataType, typename KeyType = int, typename HandleType = int>
class MaxPriorityQueue : public BinaryHeap<DataType, KeyType, HandleType>
{
  public:
    using NodeType = typename BinaryHeap<DataType, KeyType, HandleType>::NodeType;

  public:
    const HandleType &Insert(const NodeType &node)
    {
        // push to the end of the array
        this->storage_.push_back(node);
        NodeType &inserted_node = this->storage_.back();
        inserted_node.index = this->storage_.size() - 1;
        this->heap_size_ = this->storage_.size();

        // then increase key for it
        return IncreaseKey(inserted_node.index, inserted_node.key); // up
    }

    const NodeType &Maximum() const
    {
        assert(!this->storage_.empty());
        return this->storage_[0];
    }

    NodeType ExtractMax()
    {
        assert(!this->storage_.empty());

        this->Swap(0, this->storage_.size() - 1);
        this->heap_size_--;
        NodeType back = this->storage_.back();
        this->storage_.pop_back();

        MaxHeapity(1); // down

        return back;
    }

    const HandleType &IncreaseKey(const HandleType &node_handle, KeyType new_key)
    {
        if (node_handle >= this->heap_size_)
        {
            return node_handle;
        }

        NodeType &node = this->storage_[node_handle];
        if (new_key < node.key)
        {
            return node_handle;
        }

        node.key = new_key;

        int n = node.index + 1;
        while (n > 1)
        {
            int p = this->parent(n);
            if (!(this->storage_[p - 1] < this->storage_[n - 1]))
            {
                break;
            }

            this->Swap(p - 1, n - 1);
            n = p;
        }

        return this->storage_[n - 1].index; // return handle
    }

  private:
    // NOTE: `n` in below functions are all started by 1, not 0.
    void MaxHeapity(int n)
    {
        while (true)
        {
            int left = this->leftChild(n);
            int right = this->rightChild(n);

            int largest = n;

            if (left <= this->heap_size_ && this->storage_[n - 1] < this->storage_[left - 1])
            {
                largest = left;
            }
            if (right <= this->heap_size_ && this->storage_[largest - 1] < this->storage_[right - 1])
            {
                largest = right;
            }

            if (largest == n)
            {
                break;
            }

            // swap n and largest
            this->Swap(n - 1, largest - 1);
            n = largest;
        }
    }
};

} // namespace myheaps